# Wallet-Dapp 后端
desc:区块链钱包Dapp开发后端设计

## 技术栈:
- Golang的go-zero、Gorm框架
- Mysql、Redis


## 钱包模块（wallet）:
- [x] 创建钱包
- [ ] 登录钱包
- [ ] 根据助记词重置密码
- [ ] 查询钱包下的账户
- [ ] 删除钱包下的账户
- [ ] 在钱包下添加账户
> 钱包和账户是一对多关系，一个钱包可能会有多个账户。在创建钱包的时候会自动在该钱包下创建一个账户，该账户不能被删除。






## shell命令：
```shell
goctl api -o wallet.api 

goctl api go  --style go_zero -api wallet.api  -dir .

```


