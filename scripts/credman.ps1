param([string]$Op, [string]$Name)

Add-Type -TypeDefinition @'
using System;
using System.Runtime.InteropServices;
public static class Cred {
    [StructLayout(LayoutKind.Sequential, CharSet = CharSet.Unicode)]
    struct CREDENTIAL {
        public uint Flags; public uint Type; public string TargetName; public string Comment;
        public long LastWritten; public uint CredentialBlobSize; public IntPtr CredentialBlob;
        public uint Persist; public uint AttributeCount; public IntPtr Attributes;
        public string TargetAlias; public string UserName;
    }
    [DllImport("advapi32.dll", CharSet = CharSet.Unicode, SetLastError = true)]
    static extern bool CredWriteW(ref CREDENTIAL c, uint flags);
    [DllImport("advapi32.dll", CharSet = CharSet.Unicode, SetLastError = true)]
    static extern bool CredReadW(string target, uint type, uint flags, out IntPtr c);
    [DllImport("advapi32.dll", CharSet = CharSet.Unicode, SetLastError = true)]
    static extern bool CredDeleteW(string target, uint type, uint flags);
    [DllImport("advapi32.dll")]
    static extern void CredFree(IntPtr c);
    public static void Put(string target, byte[] blob) {
        var c = new CREDENTIAL { Type = 1, TargetName = target, UserName = "nlm", Persist = 2, CredentialBlobSize = (uint)blob.Length };
        c.CredentialBlob = Marshal.AllocHGlobal(blob.Length);
        Marshal.Copy(blob, 0, c.CredentialBlob, blob.Length);
        CredWriteW(ref c, 0);
        Marshal.FreeHGlobal(c.CredentialBlob);
    }
    public static byte[] Get(string target) {
        IntPtr p;
        if (!CredReadW(target, 1, 0, out p)) return null;
        var c = (CREDENTIAL)Marshal.PtrToStructure(p, typeof(CREDENTIAL));
        var blob = new byte[c.CredentialBlobSize];
        Marshal.Copy(c.CredentialBlob, blob, 0, blob.Length);
        CredFree(p);
        return blob;
    }
    public static void Del(string target) { CredDeleteW(target, 1, 0); }
}
'@

$size = 1024
function Count { $b = [Cred]::Get($Name); if ($b) { [int][Text.Encoding]::ASCII.GetString($b) } else { 0 } }

switch ($Op) {
    'set' {
        $text = [Console]::In.ReadToEnd().Trim()
        $n = [math]::Ceiling($text.Length / $size)
        for ($i = 0; $i -lt (Count); $i++) { [Cred]::Del("$Name/$i") }
        for ($i = 0; $i -lt $n; $i++) {
            $part = $text.Substring($i * $size, [math]::Min($size, $text.Length - $i * $size))
            [Cred]::Put("$Name/$i", [Text.Encoding]::ASCII.GetBytes($part))
        }
        [Cred]::Put($Name, [Text.Encoding]::ASCII.GetBytes("$n"))
    }
    'get' {
        $out = ''
        for ($i = 0; $i -lt (Count); $i++) { $out += [Text.Encoding]::ASCII.GetString([Cred]::Get("$Name/$i")) }
        [Console]::Out.Write($out)
    }
    'del' {
        for ($i = 0; $i -lt (Count); $i++) { [Cred]::Del("$Name/$i") }
        [Cred]::Del($Name)
    }
}
