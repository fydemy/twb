# Twb

An open-source shareable twibbon platform without annoying ads and sophisticated authentication built with React and Go.

## 🚀 Installation

1. Make sure you have `git` and `go` installed on your machine.
   ```shell
   git clone https://github.com/fydemy/fytwb .
   ```
2. Since we just have the backend, so start off by running the go by:
	```
	cd backend
	go mod tidy // to installed the required packages
	go run main.go // or using air
	```
## 🔍 API Spec
```
/api/v1/upload
├── GET (/:id)
	├── Req: id params
	└── Res: path
├── POST (/)
	└── Req: form-data (key: frame)
	└── Res: ID
```

## ✅ Upcoming features

- [x] Add and get frame API
- [ ] Rate limiting and caching
- [ ] Frontend ✨  

## ✨ Community & Contribute

Visit our github profile and seek for a Discord link at [fydemy.com](https://fydemy.com).
