## Nolvus Automagic Downloader

![example](./example.png)

### Background for those interested

The Nolvus modpack requires a lot of files to be downloaded from
Nexus Mods. Although possible to download the mods through manually
"clicking" the "Slow Download" button on the Nexus Mods website rendered
within the Nolvus Dashboard. This will take a substantial amount of time
to complete.

The solution of course is to just buy the premium subscription from 
Nexus Mods. But that's not why you are here ;) - nor me.

Being annoyed at manually clicking the slow download button, I decided to 
find out how to automate it. Of course, the initial thought is to use an automated
mouse clicker. This doesn't work since the window rendering the website
might not be scrolled to the correct position or some Nexus mods video ads would
popup over the button etc.

### Requirements

#### 1. Edited Nolvus Dashboard shortcut

Launch the Nolvus Dashboard with the `--remote-debugging-port=8088`.
This can be done on Windows by creating a shortcut of the executable
and then specifying the flags after the executable in the `Target` field.

Example:

```
Target: "<path to exe>" --remote-debugging-port=8088
```

#### 2. Installed GoLang

Ensure that GoLang is installed on your system. If not, follow these steps to install GoLang:

1. You can find the latest Go release [here](https://go.dev/dl/).
2. Choose the appropriate installer for your operating system (Windows, macOS, or Linux).
3. Run the downloaded installer and follow the on-screen instructions.
4. Confirm the installation by opening a terminal or command prompt and typing:
    ```powershell
    go version
    ```

5. You should see an output indicating the installed version of GoLang, confirming that the installation was successful.

### Nolvus Automagic Downloader Installation & Usage

1. **Build Nolvus Automagic Downloader**

    - Open **Windows Powershell** console

        - Press `Win + X` and select Windows PowerShell (or Windows Terminal if you have it installed).

    - Clone the Repository:

        ```powershell
        git clone git@github.com:Benehiko/nolvus-automagic-downloader.git
        cd nolvus-automagic-downloader
        ```

        - **Note:** If you aren't a skilled GitHub user but still want to download this wonderful utility, it might be easier for you to download the repository as a ZIP file (click the `< > Code` button at the top of the page and the `Download ZIP` button).

    - Build the Nolvus Automagic Downloader app:

        ```powershell
        go build .
        ```

        - **Note:** If you encounter a build error saying that there is a problem with your Go version, open the `go.mod` file in your text editor and verify your version of Go is the same or greater.

    - Don't close the console yet.

2. **Edit the Nolvus Dashboard shortcut as mentioned above**

3. **Launch the Nolvus Dashboard shortcut and start the mod download process**

    - **Note:** You might need to launch the shortcut with admin privileges (it depends on your PC settings and the Nolvus Dashboard location on your disk).

4. **Run the Nolvus Automagic Downloader**

    ```powershell
    go run .
    ```

5. **Now sit back and wait patiently for the download to finish**

    - **Note:** You should be able to use your PC normally during the download.

    - **Note:** Don't close the Console while the download is proceeding as it would close the Nolvus Automagic Downloader. If you close your Console accidentally, don't panic, just launch the Nolvus Automagic Downloader again :-)

### Run Nolvus with Wine

Nolvus is a .Net application running [Chromium Embedded Framework (CEF)](https://bitbucket.org/chromiumembedded/cef/src/master/),
specifically [CEFSharp](https://cefsharp.github.io/).

The Wine setup I found to work through [bottles](https://usebottles.com/) with the `sys-wine-9.0` runner
and the following dependencies:

```
- arial32
- times32
- courie32
- mono
- gecko
- vcredist2019
- andale32
- arialb32
- comic32
- georgi32
- impact32
- tahoma32
- trebuc32
- verdan32
- webdin32
- allfonts
- dotnet40
- dotnet45
- dotnet46
- dotnet461
- dotnet462
- dotnet472
- vcredist2022
- vcredist6
- vcredist2015
- dotnet452
- vcredist2013
- consolas
- unifont
```

Once the Nolvus Dashboard does the installation steps (with the CEF popups) after clicking 
the "Slow Download" button it might give you an error about missing fonts, specifically the
SegoeUI fonts.

Below are instructions for installing Microsoft fonts on Arch.

```console
yay -Sy ttf-win10
```

Other packages such as `ttf-ms-win10-auto` didn't work for me.
https://wiki.archlinux.org/title/Microsoft_fonts

### Misc

**A workaround if still having problems with pop-up ads.**

Running the downloads this way could re-focus the mouse on to the popup window which is very annoying if you are still using the PC.

So far I have been able to disable the rendering through CEF with additional flags but this does not prevent Nolvus from popping up a new window.

```console
--off-screen-rendering-enabled --headless --disable-gpu 
```
