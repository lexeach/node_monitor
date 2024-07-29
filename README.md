## Node monitor

This is a little tool to keep track on a set of nodes, 
and see if they keep in step or if they go out of consensus. 

## Usage
1. Copy the `config.example.toml` into a `config.toml` file and fill in the required information
2. with Docker: `docker run -d --name nodemonitor -p 8080:8080 -v <path-to-config.toml>/config.toml:/config.toml holiman/nodemonitor:latest /config.toml`
3. Access the webpage by navigating to `http://localhost:8080`

## Dashboard

It shows a neat little dashboard, where 'interesting' points of differing opinions are shown: 

![nodemon](nodemon.png)

## Metrics

It also has support for pushing metrics to `influxdb`, so you can get nice charts and 
alerts from all/any node which supports basic set of standard rpc methods. 

![](charts.png)




#command from satya for deployemnet on server

1. sudo apt install apt-transport-https ca-certificates curl software-properties-common
2.  curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
3. echo "deb [arch=amd64 signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
4. sudo apt install docker-ce docker-ce-cli containerd.io
5. sudo docker run hello-world
6. --version
7. docker run -d --name nodemonitor -p 8080:8080 -v <path-to-config.toml>/config.toml:/config.toml holiman/nodemonitor:latest /config.toml
8. docker run -d --name nodemonitor -p 8080:8080 -v ./config.toml:/config.toml holiman/nodemonitor:latest /config.toml
9. docker restart nodemonitor
10. docker stop nodemonitor
11. docker rm -f nodemonitor
