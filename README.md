## Novus Automagic Downloader

### Background for those interested

The Novus modpack requires a lot of files to be downloaded from
Nexus Mods. Although possible to download the mods through manually
"clicking" the "Slow Download" button on the Nexus Mods website rendered
within the Novus Dashboard. This will take a substantial amount of time
to complete.

The solution of course is to just buy the premium subscription from 
Nexus Mods. But that's not why you are here ;) - nor me.

Being annoyed at manually clicking the slow download button, I decided to 
find out how to automate it. Of course the inital thought is to use an automated
mouse clicker. This doesn't work since the window rendering the website
might not be scrolled to the correct position or some Nexus mods video ads would
popup over the button etc.

### Requirements

Launch the Novus Dashboard with the `--remote-debugging-port=8088`.
This can be done on Windows by creating a shortcut of the executable
and then specifying the flags after the executable in the `Target` field.

Example:

```
Target: "<path to exe>" --remote-debugging-port=8088
```

Then run the Automagic Downloader.

```
go run .
```

### Build Novus Automagic Downloader

```console
git clone git@github.com:Benehiko/novus-automagic-downloader.git
cd novus-automagic-downloader
go build .
```
