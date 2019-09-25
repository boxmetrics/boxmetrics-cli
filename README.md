# boxmetrics-cli

> This repo contains the boxmetrics cli built with Go

## ✨ Features

This CLI aim to manage the whole boxmetrics ecosystem :

- Install agent locally or on multiple target
- Uninstall agent locally or on multiple target
- Get various data from API
- Generate report from boxmetrics API
- Test boxmetrics API heatlcheck
- Test boxmetrics frontapp heatlcheck

## 🚀 Quick start

### **1. From docker image :**

```bash
docker run boxmetrics/boxmetrics-cli
```

### **2. From installation script :**

```bash
curl https://raw.githubusercontent.com/boxmetrics/boxmetrics-cli/master/scripts/install.sh | bash
```

### **3. From prebuilt binary :**

1. **Donwload a binary from Github [release page](https://github.com/boxmetrics/boxmetrics-cli/releases)**

2. **Run application**

```bash
# Made application executable
sudo chmod +x boxmetrics

# Run application
./boxmetrics --help
```

### **4. From source code :**

1. **Clone the git repository**

```bash
# cloning git repository
git clone https://github.com/boxmetrics/boxmetrics-cli
```

2. **Build application**

```bash
# go to boxmetrics-cli directory
cd boxmetrics-cli
# run helper command to build
make build
```

3. **Run application**

```bash
# Made application executable
sudo chmod +x bin/boxmetrics

# run application
./bin/boxmetrics --help
```

## 💬 Contributing

1. **Fork the git repository**

2. **Create your feature branch**

3. **Apply your changes**

4. **Run tests**

```bash
# run test
make tests

make build
# test your feature
./bin/boxmetrics -awesome feature
```

6. **Commit your changes**

7. **Push it on your fork**

8. **Create new pull request**

## 🧐 What's inside ?

```text
.
├── bin             # Project binaries
├── cmd             # Main applications for this project
├── configs         # Configuration file templates or default configs
├── docs            # Design and user documents
├── internal
│   └── pkg         # Private library code
├── scripts         # Scripts to perform various build, install, analysis, etc operations
├── test            # Additional external test apps and test data
├── Dockerfile      # Docker image
├── go.mod          # Module dependencies
├── go.sum          # Ensure dependencies integrity
├── JenkinsFile     # Jenkins pipeline
├── main.go         # Application entry point
├── Makefile        # Helpers command
├── LICENSE
└── README.md
```

## 👥 Contributors

<table width="100%">
  <tbody width="100%">
    <tr width="100%">
      <td align="center" width="33.3333%" valign="top">
        <img style="border-radius: 50%;" width="100" height="100" src="https://github.com/Laurent-PANEK.png?s=100">
        <br>
        <a href="https://github.com/Laurent-PANEK">Laurent Panek</a>
        <p>Security System Integrator</p>
      </td>
     <td align="center" width="33.3333%" valign="top">
        <img style="border-radius: 50%;"  width="100" height="100" src="https://github.com/maxencecolmant.png?s=100">
        <br>
        <a href="https://github.com/maxencecolmant">Maxence Colmant</a>
        <p>DevOps System Integrator</p>
    </td>
          <td align="center" width="33.3333%" valign="top">
        <img style="border-radius: 50%;"  width="100" height="100" src="https://github.com/abdessalamb98.png?s=100">
        <br>
        <a href="https://github.com/abdessalamb98">Abdessalam Benharira</a>
        <p>JavaScript Developer</p>
      </td>
     </tr>
  </tbody>
</table>
