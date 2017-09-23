# Connect MIDI keyboard to synthesizer using a Raspberry Pi


## Context
I recently bought an [OP-1](https://www.teenageengineering.com/products/op-1) and already had an [Arturia Minilab](https://www.arturia.com/products/hybrid-synths/minilab/overview) MIDI controller, so I wanted to connect them. That turned out not to be easy without buying an expensive interface. Then I found [this vid on YouTube](https://www.youtube.com/watch?v=crwJ56aYkw4), describing how to use a Raspberry PI as a bridge using the `aconnect` util. The only thing missing, was that I don't want to take out my PC or screen + keyboard during a rehearsal session. Ideally, plugging in the PI and connecting my keyboard and my OP-1 would instantly connect them.

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

