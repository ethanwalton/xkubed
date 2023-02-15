# xkubed

Open source k8 security agent

## Description

Kubernetes open source container agent that is able to scan for vulnerabilties inside of your kubernetes environments. This tool will scan incrementaly over all resources located in your cluster and develop a reported output based on vulnerabilties that occur in your cluster. 

## Getting Started

### Dependencies

* Go
* Kubernetes CLI
* Docker
* Helm 3.0+

### Installing

```
kubectl apply -f /xkubed.yaml
```

### Executing program

```
cd server

docker build -t SERVER_TAG_NAME .

docker push SERVER_TAG_NAME
```

## Help

Any advice for common problems or issues.

## Authors

Ethan Walton - `Senior Security Engineer @ Ripple`
github - @ethanwalton

## Version History

* 0.1
    * Initial Release

## License

This project is licensed under the MIT License - see the LICENSE file for details

## Acknowledgments

Inspiration, code snippets, etc.
