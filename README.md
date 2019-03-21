# Red Alert 2 One Man Deathmatch

## Description

One Man Deathmatch is a gamemode for Red Alert 2.
All teams can only make one of each structure, defense, and unit

Some additional changes also exist:

1. Yuri
    1. Yuri cannot build the vehicle version of the slave miner (as that would allow you to get two slave miners)
    1. Yuri cannot build the Psychic tower, as we felt that would be overpowered.
    1. Yuri cannot build the cloning vats as it would break the spirit of the game
1. Allies
    1. America: paradrop only drops one GI.
    1. Battle fortress is disabled as it was too strong in our tests
1. Soviets
    1. No changes yet!

## Installation

The file "spawner.xdp" can be used without the launcher with CnCNet launcher. you can merely overwrite your old file to make the mod work. (I would recommend backing your old one up!)
If you don't have CnCNet you should be able to rename spawner.xdp to rulesmd.ini but I have not tested this

I also have a launcher which makes life a bit easier to not overwrite the gamedata and play the mod. All you need to do to run it is to run it with the spawner.xdp in the same place you run the program from. This has no logic specific to my mod and could be used with any spawner.xdp

My launcher implementation assumes that you have a 64 bit windows system and that you have CnCNet Installed with the origin version of the game. I might change this to be a more robust implementation if it proves necessary. (it currently looks for `C:\Program Files (x86)\Origin Games\Command and Conquer Red Alert II`)

## Building from source

1. You can download this repository to a directory.
1. You must download rsrc (go get github.com/akavel/rsrc)
1. then run `rsrc -manifest RA2OMD.exe.manifest`
1. Then the resulting exe should work