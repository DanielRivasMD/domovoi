####################################################################################################

_default:
  @just --list

####################################################################################################

# print justfile
[group('just')]
@show:
  bat .justfile --language make

####################################################################################################

# edit justfile
[group('just')]
@edit:
  micro .justfile

####################################################################################################

# compile roadmap action items
[group('dev')]
@roadmap:
  { echo '\n=================================================='; todor -s; echo '=================================================='; } >> ROADMAP.txt

####################################################################################################
# import
####################################################################################################

# config
import '.just/golib.conf'

####################################################################################################
# jobs
####################################################################################################

# install locally
[group('go')]
build: 
  @echo "\n\033[1;33mBuilding\033[0;37m...\n=================================================="
  go build *.go

####################################################################################################

# watch changes
[group('go')]
watch:
  watchexec --clear --watch captureExecCmd.go --watch changeDir.go --watch createDir.go --watch createFile.go --watch currentDir.go --watch dirExist.go --watch domovoi.go --watch ensureDirExist.go --watch execCmd.go --watch fileExist.go --watch findHome.go --watch isHidden.go --watch lineBreak.go --watch listDirs.go --watch listFiles.go --watch recallDir.go --watch walk.go -- 'just build'

####################################################################################################
