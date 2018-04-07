#Minerstats

##Usage
Run the fully built program while running a miner. It will return a JSON object with the running miners and hashrates.

##Build notes:

####Compiling
To compile this, you need to put the source in the following location:

$GOROOT/src/bitbucket.org/minerstats

Otherwise you will get a $GOROOT issue.

####Packing
The makefile has a commented out command to pack the fully built executable. This requires UPX to be installed on the machine on which you are building. UPX reduces the binary size by 70%, but may, however, trigger some antivirus. Upon release, this will be uncommented and all builds will be packed.
