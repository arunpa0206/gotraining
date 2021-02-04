Installion in windows

step 1. install chocolatey using the following command:
     open windows powershell with admin access
     Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))


step 2. install make using the following command:
    choco install make
    make --version

step 3. install-git using following command:
choco install git
Restart the terminal 

Step 4 . Install golang using the following command:
        choco install golang


step 5. setup workspace 
        echo %GOPATH%
        cd  %GOPATH%
        mkdir bin pkg src
        cd src
        mkdir github.com
        cd github.com
        mkdir $git repo (replace with your git repo name)
        cd $repo

step 6. Run the following make target
        make install-windows

step 7. Test your environment
        make test

step 8 get the libriraries
        make get

Installion in MAC

step 1. Install homebrew uisng following command:
         ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"

step 2. Install make using following command
        brew install make
        make --version

step 3. Install git using following command
        brew install git
        git --version

step 4. Run the following make target
        make install-mac

step 5. setup workspace
        code ~/.bash_profile

5.1 Add the following to this file:
        export GOPATH=$workdirectory/go (replace workdirectory with path )
        export GOBIN=$GOPATH/bin
        export PATH=$PATH:$GOPATH:$GOBIN
        export PATH=$PATH:/usr/local/sbin

5.2 Cretae the recommended directories
        cd $GOPATH
        mkdir bin pkg src
        cd src
        mkdir github.com
        cd github.com
        mkdir $git repo (replace with your git repo name)
        cd $repo
        git clone https://github.com/arunpa0206/gotraining.git

step 7. test your environment:
        make test

step 8 get the libriraries
        make get