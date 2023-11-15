<p align="center">
  <img width="180" src="https://yvling-typora-image-1257337367.cos.ap-nanjing.myqcloud.com/typora/dog.png" alt="ChatGPT">
  <h1 align="center">Server Dog</h1>
  <p align="center">简单易用的服务器探针</p>
</p>



## 项目简介

ServerDog是一个使用Golang实现的轻量服务器探针，具有简单、易用等特点，只需一键即可连接n台服务器，实时显示各服务器的运行状态。



## 使用方法

从 [Release](https://github.com/Fangnan700/ServerDog/releases) 处下载最新的压缩包，在服务器上解压得到 `server_dog` ：

```shell
tar -xvJf server_dog_linux_1.0.0_amd64.tar.xz
```

切换至 `root` 用户：

```shell
sudo -s
```

授予可执行文件权限：

```shell
chmod +x server_dog
```

运行程序：

```shell
./server_dog -P 7776 -p 7777
```

参数说明：

- -P：web面板端口，使用浏览器访问查看服务器状态
- -p：websocket端口，节点间通过此端口获取数据
- 不带参数时web面板默认使用 `7776` 端口，websocket默认使用 `7777` 端口

访问面板：

访问服务器 `7776` 端口即可打开面板查看服务器运行状态。

![image-20231115165700696](https://yvling-typora-image-1257337367.cos.ap-nanjing.myqcloud.com/typora/image-20231115165700696.png)

连接其他机器：

在待连接服务器上按上述步骤运行程序，同时放行websocket端口；

在面板右上角输入待连接服务器的ip+websocket端口，即可连接。

![image-20231115170016841](https://yvling-typora-image-1257337367.cos.ap-nanjing.myqcloud.com/typora/image-20231115170016841.png)





