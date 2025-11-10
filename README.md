# go-tintlog

A tiny, zero-dependency logging helper for Go that prints **clean, readable, color-tinted**
messages to terminals while keeping **file output uncolored** for machines and post-processing.

## Why
- Human eyes get quick signal from color; machines need plain text/JSONL.
- Truecolor (24-bit) ANSI done simply, with a small curated palette.
- Minimal surface area; drop-in and forget.

## What it provides
- Per-line tinting with optional bold, designed for real terminals.
- A clear, editor-friendly color palette (hex strings) with base, Bright, and Dim variants.
- A lightweight colorizer registry for consistent styles across your app.
- Utilities for pretty/compact value rendering and safe argument sanitization.
- Dockerfile and docker-compose.yml files to test colors with docker-compose up
