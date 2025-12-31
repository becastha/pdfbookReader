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
│   ├── handlers/                      # HTTP layer 
│   │   ├── document_handler.go        # CRUD for documents
│   │   ├── page_handler.go            # Get pages, content
│   │   ├── content_handler.go         # Get content blocks
│   │   ├── upload_handler.go          # PDF upload endpoint
│   │   ├── search_handler.go          # Full-text search
│   │   └── health_handler.go
│   │
│   ├── services/                      # Business logic
│   │   ├── document_service.go        # Document CRUD operations
│   │   ├── extraction_service.go      # Orchestrates extraction pipeline
│   │   ├── search_service.go          # Search logic
│   │   └── thumbnail_service.go
│   │
│   ├── extraction/                    # PDF extraction layer (NEW)
│   │   ├── extractor.go               # Main interface
│   │   ├── pdfcpu.go                  # pdfcpu wrapper (metadata, TOC)
│   │   ├── poppler.go                 # poppler wrapper (text extraction)
│   │   ├── content_parser.go          # Parse text into blocks
│   │   └── table_detector.go          # Table detection logic
│   │
│   ├── repository/                    # Database layer 
│   │   ├── document_repo.go
│   │   ├── page_repo.go
│   │   ├── content_block_repo.go
│   │   ├── outline_repo.go
│   │   ├── image_repo.go
│   │   └── table_repo.go
│   │
│   ├── models/                        # schema
│   │   ├── document.go
│   │   ├── page.go
│   │   ├── content_block.go
│   │   ├── document_outline.go
│   │   ├── image.go
│   │   └── table.go
│   ├── routes/                        # Routing layer
│   │   ├── router.go                  # Main router setup
│   │   ├── document_routes.go
│   │   ├── page_routes.go
│   │   ├── content_routes.go
│   │   ├── search_routes.go
│   │   └── health_routes.go
│   │
│   │
│   ├── worker/                        # Background processing 
│   │   ├── processor.go               # Job processor
│   │   └── jobs/
│   │       ├── extract_pdf.go         # PDF extraction job
│   │       └── generate_thumbnail.go
│   │
│   ├── database/
│   │   ├── postgres.go
│   │   └── migrations/
│   │       ├── 001_documents.sql
│   │       ├── 002_pages.sql
│   │       ├── 003_content_blocks.sql
│   │       ├── 004_document_outline.sql
│   │       ├── 005_images.sql
│   │       └── 006_tables.sql
│   │
│   └── config/
│       └── config.go
│
├── pkg/
│   └── utils/
│       ├── file_validator.go
│       ├── hash.go                    # File hashing for dedup
│       └── logger.go
│
├── storage/
│   ├── pdfs/                          # Original PDFs
│   ├── thumbnails/
│   └── images/                        # Extracted images 
│
├── go.mod
├── go.sum
├── Dockerfile
└── docker-compose.yml
```
