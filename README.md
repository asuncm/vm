# vm虚拟世界
项目中的公用组件项目

#### Git提交规范

```base
类型        描述
feat        新增 feature 新功能、新特性
fix         修复bug
docs	    仅仅修改了文档，比如 README...
style	    修改了空格、格式缩进、逗号等，不改变代码逻辑
refactor    代码重构，没有加新功能或者修复bug
perf	    优化相关，比如提升性能、体验
test	    测试用例，比如单元测试、集成测试等
chore	    改变构建流程、或者增加依赖库、工具等
build       影响项目构建或依赖项修改
ci          持续集成相关文件修改
revert	    回滚到上一个版本
release     发布新版本
workflow    工作流相关文件修改
```

#### 项目环境配置说明

```text
 1. 项目根目录config文件夹存放所有配置信息
 2. 项目启动参数 mode=xxx 为配置文件的前缀名称
 3. 配置文件命名以 xxx.config.yaml为准
 4. env.config.yaml为参照模板，文件命名格式务必与环境参数mode的值保持一致
 
 示例：本地开发环境
 mode=local
 配置文件：local.config.yaml
```

#### API接口状态码

```bash
1. 项目API接口返回状态码，详情查询service模块README.md文件说明
```

#### 项目结构说明

```bash
|-config
| |-env.config.yaml   环境配置文件示例
| └─*.config.yaml     程序运行环境配置文件
|-service 
|-auth
|-user
|-authentication
|-design
|-doc
|-flow
|-form
|-task
|-web
|-website
|-teach
```