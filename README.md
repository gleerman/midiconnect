# Connect USB MIDI keyboard to synthesizer using a Raspberry Pi


## Context
I recently bought an [OP-1](https://www.teenageengineering.com/products/op-1) and already had an [Arturia Minilab](https://www.arturia.com/products/hybrid-synths/minilab/overview) USB MIDI controller, so I wanted to connect them. That turned out not to be easy without buying an expensive interface. Then I found [this vid on YouTube](https://www.youtube.com/watch?v=crwJ56aYkw4), describing how to use a Raspberry PI as a bridge using the `aconnect` util. The only thing missing, was that I don't want to take out my PC or screen + keyboard during a rehearsal session. Ideally, plugging in the PI and connecting my keyboard and my OP-1 would instantly connect them.

This script can be added as a cron job and does exactly that. This is what my PI says now:
```
$ sudo crontab -l
@reboot /path/to/midiconnect "Arturia MINILAB" "OP-1 Midi Device"
*/1 * * * * /path/to/midiconnect "Arturia MINILAB" "OP-1 Midi Device"
```

The script was written in Go and was my first project in that language.

# Install
Take the binary from the `dist` folder, move it the Rasperry Pi and execute it.

# Usage
```
midiconnect <src> <dest>
   src    The name of the keyboard MIDI interface
   dest   The name of the MIDI interface to be driven by the keyboard
   ```

# FAQ
##### How do I put this on my Raspberry Pi?
You could either use SCP to copy the binary to your Raspberry Pi, or download it directly if your Pi has internet access. The latter is performed by entering the following commands on your Pi:
```
cd /desired/path/
sudo apt-get install wget
wget https://github.com/gleerman/midiconnect/raw/master/dist/linux_armv5/midiconnect
```

##### How do I make it so that the script is performed frequently in order to automatically connect the USB devices?
Perform the following command on your Raspberry Pi:
```
sudo crontab -e
```
Crontab is an automated job that performs specific actions on a set schedule.
If this is the first time you edit the cron jobs, select your favourite editor. Then, at the bottom of the file that opened, enter the following lines
```
@reboot /path/to/midiconnect "<NameOfController>" "<NameOfSynth>"
*/1 * * * * /path/to/midiconnect "<NameOfController>" "<NameOfSynth>"
```
Save and close the file. If editing in Gnu nano, ctrl+O to save and ctrl+X to close.

##### How do I know the names of my controller and my synth to call `midiconnect`with?
Plug both devices to your Raspberry Pi and perform the following command:
```
aconnect -i
```
You'll see both in the list. If you're not sure, compare the list with the output if the USB devices are not connected. The name is between quotes on the lines that start with `client`.

