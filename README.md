# smartparks-connect-web
Web application for connecting to Smart Parks OpenCollar sensors.

## QuickStart
1. Compile front-end project： ```cd public && npm install && npm run build && cd ../```
2. Start server ```go run main.go```

## Development

### Backend
To start development server for backend you need to run ```go run main.go```

### Frontend
To work on frontend the best way is to run development version of frontend.
```cd public && npm run dev```. Frontend will be running on http://localhost:5173/

To build frontend for production run ```cd public && npm install && npm run build```