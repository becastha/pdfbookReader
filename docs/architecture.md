# Architecture for this project

## Frontend

```
frontend/
├── public/
│   └── index.html
├── src/
│   ├── components/
│   │   ├── PDFViewer/
│   │   │   ├── PDFViewer.jsx          # Main PDF display component
│   │   │   ├── PageControls.jsx       # Next/Prev buttons, page input
│   │   │   ├── BookView.jsx           # Book-like layout
│   │   │   └── SinglePageView.jsx     # Traditional view
│   │   ├── Layout/
│   │   │   ├── Header.jsx             # Top bar with controls
│   │   │   ├── Sidebar.jsx            # Thumbnail navigation
│   │   │   └── FullscreenWrapper.jsx  # Fullscreen mode handler
│   │   ├── Upload/
│   │   │   └── PDFUploader.jsx        # Drag & drop upload
│   │   └── UI/
│   │       ├── ThemeToggle.jsx        # Dark/Light switch
│   │       ├── ZoomControls.jsx       # Zoom in/out
│   │       └── ProgressBar.jsx        # Reading progress
│   ├── hooks/
│   │   ├── usePDF.js                  # PDF loading & state
│   │   ├── useTheme.js                # Theme management
│   │   └── useReadingProgress.js      # Save/load progress
│   ├── services/
│   │   └── api.js                     # All backend API calls
│   ├── context/
│   │   └── AppContext.jsx             # Global state
│   ├── styles/
│   │   ├── themes.css                 # Dark/Light mode CSS
│   │   └── animations.css             # Page flip animations
│   ├── App.jsx
│   └── main.jsx
├── package.json
└── vite.config.js
```

## Backend 

```
backend/
├── cmd/
│   └── main.go                        # Application entry point
├── internal/
│   ├── handlers/                      # HTTP request handlers
│   │   ├── pdf_handler.go             # PDF upload/serve endpoints
│   │   ├── progress_handler.go        # Reading progress
│   │   └── health_handler.go          # Health check
│   ├── services/                      # Business logic
│   │   ├── pdf_service.go             # PDF processing
│   │   ├── storage_service.go         # File storage operations
│   │   └── thumbnail_service.go       # Generate thumbnails
│   ├── models/                        # Data structures
│   │   ├── pdf.go                     # PDF model
│   │   └── progress.go                # Reading progress model
│   ├── database/                      # Database layer
│   │   ├── postgres.go                # PostgreSQL connection
│   │   └── migrations/                # Database schema
│   │       └── 001_init.sql
│   └── config/
│       └── config.go                  # Configuration management
├── pkg/
│   └── utils/                         # Shared utilities
│       ├── file_validator.go          # Validate PDF files
│       └── logger.go                  # Logging
├── storage/                           # Local file storage (dev)
│   ├── pdfs/
│   └── thumbnails/
├── go.mod
├── go.sum
├── Dockerfile
└── docker-compose.yml
```
