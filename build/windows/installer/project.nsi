Unicode true

## default privilege level
## !define REQUEST_EXECUTION_LEVEL "admin"

!include "wails_tools.nsh"

# The version information for this two must consist of 4 parts
VIProductVersion "${INFO_PRODUCTVERSION}.0"
VIFileVersion    "${INFO_PRODUCTVERSION}.0"

VIAddVersionKey "CompanyName"     "${INFO_COMPANYNAME}"
VIAddVersionKey "FileDescription" "${INFO_PRODUCTNAME} Installer"
VIAddVersionKey "ProductVersion"  "${INFO_PRODUCTVERSION}"
VIAddVersionKey "FileVersion"     "${INFO_PRODUCTVERSION}"
VIAddVersionKey "LegalCopyright"  "${INFO_COPYRIGHT}"
VIAddVersionKey "ProductName"     "${INFO_PRODUCTNAME}"

!include "MUI.nsh"

!define MUI_ICON "..\icon.ico"
!define MUI_UNICON "..\icon.ico"
;!define MUI_WELCOMEFINISHPAGE_BITMAP "..\leftimage.bmp" # !!!Include this to add a bitmap on the left side of the Welcome Page. Must be a size of 164x314
;!define MUI_UNWELCOMEFINISHPAGE_BITMAP "..\leftimage.bmp" # !!!Include this to add a bitmap on the left side of the Welcome Page. Must be a size of 164x314
!define MUI_FINISHPAGE_NOAUTOCLOSE # Wait on the INSTFILES page so the user can take a look into the details of the installation steps
!define MUI_ABORTWARNING # This will warn the user if they exit from the installer.

!define MUI_WELCOMEPAGE_TITLE "${INFO_PRODUCTNAME} ${INFO_PRODUCTVERSION}"
#!define MUI_WELCOMEPAGE_TEXT "${INFO_PRODUCTVERSION}$\r$\n${mui.WelcomePage.Text}"
#!define MUI_PAGE_HEADER_SUBTEXT "Hello, world ${INFO_PRODUCTVERSION}"

!insertmacro MUI_PAGE_WELCOME # Welcome to the installer page.
# !insertmacro MUI_PAGE_LICENSE "resources\eula.txt" # Adds a EULA page to the installer
!insertmacro MUI_PAGE_DIRECTORY # In which folder install page.
!insertmacro MUI_PAGE_INSTFILES # Installing page.
!insertmacro MUI_PAGE_FINISH # 确认结束

!insertmacro MUI_UNPAGE_WELCOME
!insertmacro MUI_UNPAGE_CONFIRM # 确认卸载
!insertmacro MUI_UNPAGE_INSTFILES # Uinstalling page

# 添加多语言的宏
  !insertmacro MUI_LANGUAGE "English"
  !insertmacro MUI_LANGUAGE "French"
  !insertmacro MUI_LANGUAGE "German"
  !insertmacro MUI_LANGUAGE "Spanish"
  !insertmacro MUI_LANGUAGE "SpanishInternational"
  !insertmacro MUI_LANGUAGE "SimpChinese"
  !insertmacro MUI_LANGUAGE "TradChinese"
  !insertmacro MUI_LANGUAGE "Japanese"
  !insertmacro MUI_LANGUAGE "Korean"
  !insertmacro MUI_LANGUAGE "Italian"
  !insertmacro MUI_LANGUAGE "Dutch"
  !insertmacro MUI_LANGUAGE "Danish"
  !insertmacro MUI_LANGUAGE "Swedish"
  !insertmacro MUI_LANGUAGE "Norwegian"
  !insertmacro MUI_LANGUAGE "NorwegianNynorsk"
  !insertmacro MUI_LANGUAGE "Finnish"
  !insertmacro MUI_LANGUAGE "Greek"
  !insertmacro MUI_LANGUAGE "Russian"
  !insertmacro MUI_LANGUAGE "Portuguese"
  !insertmacro MUI_LANGUAGE "PortugueseBR"
  !insertmacro MUI_LANGUAGE "Polish"
  !insertmacro MUI_LANGUAGE "Ukrainian"
  !insertmacro MUI_LANGUAGE "Czech"
  !insertmacro MUI_LANGUAGE "Slovak"
  !insertmacro MUI_LANGUAGE "Croatian"
  !insertmacro MUI_LANGUAGE "Bulgarian"
  !insertmacro MUI_LANGUAGE "Hungarian"
  !insertmacro MUI_LANGUAGE "Thai"
  !insertmacro MUI_LANGUAGE "Romanian"
  !insertmacro MUI_LANGUAGE "Latvian"
  !insertmacro MUI_LANGUAGE "Macedonian"
  !insertmacro MUI_LANGUAGE "Estonian"
  !insertmacro MUI_LANGUAGE "Turkish"
  !insertmacro MUI_LANGUAGE "Lithuanian"
  !insertmacro MUI_LANGUAGE "Slovenian"
  !insertmacro MUI_LANGUAGE "Serbian"
  !insertmacro MUI_LANGUAGE "SerbianLatin"
  !insertmacro MUI_LANGUAGE "Arabic"
  !insertmacro MUI_LANGUAGE "Farsi"
  !insertmacro MUI_LANGUAGE "Hebrew"
  !insertmacro MUI_LANGUAGE "Indonesian"
  !insertmacro MUI_LANGUAGE "Mongolian"
  !insertmacro MUI_LANGUAGE "Luxembourgish"
  !insertmacro MUI_LANGUAGE "Albanian"
  !insertmacro MUI_LANGUAGE "Breton"
  !insertmacro MUI_LANGUAGE "Belarusian"
  !insertmacro MUI_LANGUAGE "Icelandic"
  !insertmacro MUI_LANGUAGE "Malay"
  !insertmacro MUI_LANGUAGE "Bosnian"
  !insertmacro MUI_LANGUAGE "Kurdish"
  !insertmacro MUI_LANGUAGE "Irish"
  !insertmacro MUI_LANGUAGE "Uzbek"
  !insertmacro MUI_LANGUAGE "Galician"
  !insertmacro MUI_LANGUAGE "Afrikaans"
  !insertmacro MUI_LANGUAGE "Catalan"
  !insertmacro MUI_LANGUAGE "Esperanto"
  !insertmacro MUI_LANGUAGE "Asturian"
  !insertmacro MUI_LANGUAGE "Basque"
  !insertmacro MUI_LANGUAGE "Pashto"
  !insertmacro MUI_LANGUAGE "ScotsGaelic"
  !insertmacro MUI_LANGUAGE "Georgian"
  !insertmacro MUI_LANGUAGE "Vietnamese"
  !insertmacro MUI_LANGUAGE "Welsh"
  !insertmacro MUI_LANGUAGE "Armenian"
  !insertmacro MUI_LANGUAGE "Corsican"
  !insertmacro MUI_LANGUAGE "Tatar"
  !insertmacro MUI_LANGUAGE "Hindi"

## The following two statements can be used to sign the installer and the uninstaller. The path to the binaries are provided in %1
#!uninstfinalize 'signtool --file "%1"'
#!finalize 'signtool --file "%1"'

Name "${INFO_PRODUCTNAME}"
!define RegName "CSGO Demo Downloader" # english name for register

InstallDirRegKey HKLM "Software\${INFO_COMPANYNAME}\${RegName}" "Install_Dir" # Memorize previous installed directory
; InstallDir "$PROGRAMFILES64\${INFO_COMPANYNAME}\${INFO_PRODUCTNAME}"
OutFile "..\..\bin\${INFO_PROJECTNAME}-${INFO_ProductVersion}-${ARCH}-installer.exe" # Name of the installer's file. 架构 -${ARCH}

ShowInstDetails hide # hide or show ==> This will always show the installation details.
ShowUninstDetails hide # hide or show

AutoCloseWindow true

Function readInstallDir
   # set 32/64bit reg
   ${If} ${RunningX64}
      SetRegView 64
   ${EndIf}

   # read InstallDirRegKey
   ReadRegStr $1 HKLM "Software\${INFO_COMPANYNAME}\${RegName}" "Install_Dir"
   ${If} $1 != ""
      StrCpy $INSTDIR $1
   ${Else}
      StrCpy $INSTDIR "$PROGRAMFILES64\${INFO_COMPANYNAME}\${INFO_PRODUCTNAME}"
   ${EndIf}
FunctionEnd

Function .onInit
   System::Call 'SHCore::SetProcessDpiAwareness(i 2)i.R0'
   !insertmacro wails.checkArchitecture

   call readInstallDir
   ; inetc::get /header "Content-Type:application/json" 'http://127.0.0.1:55608/quit' /end
   ; ExecShell open 'http://127.0.0.1:55608/quit'
FunctionEnd

Function un.onInit
   System::Call 'SHCore::SetProcessDpiAwareness(i 2)i.R0'
   ; inetc::get /header "Content-Type:application/json" 'http://127.0.0.1:55608/quit' /end
   ; ExecShell open 'http://127.0.0.1:55608/quit'
FunctionEnd

Function .onInstSuccess
   ExecShell "" "$INSTDIR\${INFO_PRODUCTNAME}.exe"
FunctionEnd

Section
    ;inetc::get /header "Content-Type:application/json" 'http://127.0.0.1:55608/quit' /end # 退出正在运行的进程
    ;Sleep 300
    !insertmacro wails.webview2runtime

    SetOutPath $INSTDIR

    !insertmacro wails.files

    WriteRegStr HKLM "Software\${INFO_COMPANYNAME}\${RegName}" "Install_Dir" $INSTDIR

    CreateShortcut "$SMPROGRAMS\${INFO_PRODUCTNAME}.lnk" "$INSTDIR\${PRODUCT_EXECUTABLE}"
    CreateShortCut "$DESKTOP\${INFO_PRODUCTNAME}.lnk" "$INSTDIR\${PRODUCT_EXECUTABLE}"

    !insertmacro wails.writeUninstaller
SectionEnd

Section "uninstall"
    inetc::get /header "Content-Type:application/json" 'http://127.0.0.1:55608/quit' /end # 退出正在运行的进程
    Sleep 300

    RMDir /r "$AppData\${PRODUCT_EXECUTABLE}" # Remove the WebView2 DataPath

    # RMDir /r $INSTDIR

    # 只删除程序，避免删除备份文件
    Delete "$INSTDIR\${PRODUCT_EXECUTABLE}"

    Delete "$SMPROGRAMS\${INFO_PRODUCTNAME}.lnk"
    Delete "$DESKTOP\${INFO_PRODUCTNAME}.lnk"

    # delete previous install directory
    ${If} ${RunningX64}
      SetRegView 64
    ${EndIf}
    DeleteRegKey HKLM "Software\${INFO_COMPANYNAME}\${RegName}"
    # WriteRegStr HKLM "Software\${INFO_COMPANYNAME}\${RegName}" "Install_Dir" ""

    !insertmacro wails.deleteUninstaller
SectionEnd
