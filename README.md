# `kubeswitch`: Lightweight alternative for switching between Kubernetes clusters

---

`kubeswitch` is a tool that allows you to switch between Kubernetes clusters in a simple and fast way. It is a lightweight alternative to [kubectx](https://github.com/ahmetb/kubectx/) and [kubeswitch](https://github.com/danielfoehrKn/kubeswitch/)

![](https://s11.gifyu.com/images/ScuIM.gif)

Actually, I build this project for myself. I wanted to switch between Kubernetes clusters in a simple way. I didn't need any other features. So I build this tool.

If you need more features, you should consider using [kubectx](https://github.com/ahmetb/kubectx/) or [kubeswitch](https://github.com/danielfoehrKn/kubeswitch/) instead of this tool. Both are great tools with more features.

## Installation

##### Go download

```sh
go download github.com/anilsenay/kubeswitch
```

## Usage

##### List all contexts and switch between them

```sh
kubeswitch
```

##### Get only current contexts

```sh
kubeswitch [current | --current]
```

## Configuration

`kubeswitch` uses the `KUBECONFIG` environment variable to find the Kubernetes configuration file. If you don't set the `KUBECONFIG` environment variable, `kubeswitch` will use the default Kubernetes configuration file (`~/.kube/config`).

Also you can set it manually with `--config` flag.

```sh
kubeswitch --config=/path/to/kubeconfig
```
