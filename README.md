# go-micro-cn

这是一个 go-micro 项目，基于 micro 脚手架搭建而成，主要参考了 https://github.com/real-jacket/tutorials/tree/master/microservice-in-micro 这个教程进行开发。

项目目前的基本架构如下图（未完成）：

![](https://raw.githubusercontent.com/micro-in-cn/tutorials/master/microservice-in-micro/docs/part2_auth_layer_view.png)

## 背景

关于 go-micro ，是用 go 语言开发的微服务框架，它提供了一种开发模式去对微服务开发与治理。在此之前，我们得先弄明白微服务的基本概念。

### 什么是微服务

根据 wiki 百科 里的描述：微服务 (Microservices) 是一种软件架构风格，它是以专注于单一责任与功能的小型功能区块 (Small Building Blocks) 为基础，利用模块化的方式组合出复杂的大型应用程序，各功能区块使用与语言无关 (Language-Independent/Language agnostic) 的 API 集相互通信。

简单理解，就是微服务是一种软件架构体系，它将后端服务，根据业务进行拆分，使每个服务变成更小的独立单元，服务之间彼此互不影响，通过语言无关的通信机制进行沟通（比如 HTTP）。

![基本架构图](https://cdn.nlark.com/yuque/0/2020/png/244078/1599117343966-a752229e-18a4-461e-a40a-e92f3dac48be.png#align=left&display=inline&height=370&margin=%5Bobject%20Object%5D&name=Microservice_Architecture.png&originHeight=370&originWidth=539&size=68364&status=done&style=shadow&width=539)

### 微服务的独立性

微服务的规划与单体式应用十分不同，微服务中每个服务都需要避免与其他服务有所牵连，且都要能够自主，并在其他服务发生错误时不受干扰。意思就是每个服务就相当于独立的应用，都可以有自己的数据库，服务资源等等。

#### 数据库

从微服务的理念来看，微服务的数据应该是自己进行管理，无论数据是怎么管理，服务彼此之间的数据应该是隔离的。

- 每个服务都各有一个数据库，同属性的服务可共享同个数据库。
- 所有服务都共享同个数据库，但是不同表格，并且不会跨域访问。
- 每个服务都有自己的数据库，就算是同属性的服务也是，数据库并不会共享。

数据库并不会只存放该服务的资料，而是“**该服务所会用到的所有资料**”。此举是为了避免服务之间的相依性，避免服务之间产生耦合。

#### 沟通与时间广播

微服务中最重要的就是每个服务的独立与自主，因此服务与服务之间也**不应该**有所沟通。倘若真有沟通，也应采用**异步**沟通的方式来避免紧密的相依性问题。故要达此目的，可以采用以下两种方式：

- 沟通与事件广播：这可以让你在服务集群中广播事件，并且在每个服务中监听这些事件并作处理，这令服务之间没有紧密的相依性，而这些发生的事件都会被保存在事件存储中心里。这意味着当微服务重新上线、部署时可以重播（Replay）所有的事件。这也造就了微服务的数据库随时都可以被删除、摧毁，且不需要从其他服务中获取资料。
- 消息队列：这令你能够在服务集群中广播消息，并传递到每个服务中。与事件存储中心近乎相似，但有所不同的是：消息队列并**不会**保存事件。

#### 服务探索

随着服务的不断增加，服务之间的通信可能变得复杂，这就需要进行服务管理，产生了服务中心这样的角色。单个微服务在上线的时候，会向服务探索中心（如：Consul）注册自己的 IP 位置、服务内容，如此一来就不需要向每个微服务表明自己的 IP 位置，也就不用替每个微服务单独设置。当服务需要调用另一个服务的时候，会去询问服务探索中心该服务的 IP 位置为何，得到位置后即可直接向目标服务调用。

### 服务间的通信

关于微服务之间的通信，最主要的特性应该是跟语言平台无关，这个说的就是 http 协议了。go-micro 中服务之间的通信方式主要采用 gRPC。关于 grpc 的介绍参考[官方介绍](https://grpc.io/docs/what-is-grpc/introduction/)。

基本理念就是 通过通用的数据结构进行跨语言的交流。

![grpc](https://cdn.nlark.com/yuque/0/2020/svg/244078/1599122587686-a5ac879f-2f21-490a-817d-83a57afc6703.svg#align=left&display=inline&height=327&margin=%5Bobject%20Object%5D&name=landing-2.svg&originHeight=327&originWidth=552&size=114389&status=done&style=none&width=552)<br />

### 了解微服务

- 这里有一个网站可以了解微服务的相关理念与实践 https://microservices.io/ 。
- 这里有一个 go-micro 的中文教程 https://microhq.cn/index-cn
- 了解学习 gRPC 的基本概念，并知道各个语言的实现 https://grpc.io/docs/what-is-grpc/introduction/
