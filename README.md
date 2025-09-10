# Twb

An open-source shareable twibbon platform without annoying ads and sophisticated authentication built with React and Go.

## 🚀 Installation

1. Make sure you have `git`, `go`, and `pnpm` installed on your machine.
   ```shell
   git clone https://github.com/fydemy/fytwb .
   ```
2. Start off by running the backend server:
   ```
   cd backend
   go mod tidy // to installed the required packages
   cp .env.example .env
   go run main.go // or using air
   ```
3. Run the frontend by
   ```
   cd frontend
   pnpm install
   cp .env.example .env.local
   pnpm dev
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
- [x] Frontend ✨
- [ ] Image cropping/edit
- [ ] Rate limiting and caching
- [ ] Manageable data expiration
- [ ] How? without database at all!

## ✨ Community & Contribute

Visit our github profile and seek for a Discord link at [fydemy.com](https://fydemy.com).
