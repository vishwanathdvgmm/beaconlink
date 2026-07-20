# ==============================================================================
# BeaconLink Makefile
#
# Primary development entry point.
#
# Most targets delegate to scripts in ./scripts to keep this Makefile concise
# and platform-independent.
# ==============================================================================

SHELL := /usr/bin/env bash

.DEFAULT_GOAL := help

# ------------------------------------------------------------------------------
# Directories
# ------------------------------------------------------------------------------

SCRIPT_DIR := scripts

# ------------------------------------------------------------------------------
# Helpers
# ------------------------------------------------------------------------------

.PHONY: help

help:
	@echo ""
	@echo "BeaconLink Development Commands"
	@echo "==============================="
	@echo ""
	@echo "Setup"
	@echo "  make bootstrap      Prepare development environment"
	@echo ""
	@echo "Development"
	@echo "  make generate       Run code generation"
	@echo "  make format         Format source code"
	@echo "  make lint           Run linters"
	@echo "  make test           Run all tests"
	@echo "  make build          Build all binaries"
	@echo "  make dev            Start development environment"
	@echo "  make clean          Remove generated artifacts"
	@echo ""
	@echo "Release"
	@echo "  make release        Build release artifacts"
	@echo ""
	@echo "Go Utilities"
	@echo "  make tidy           Run go mod tidy"
	@echo "  make download       Download Go modules"
	@echo "  make verify         Verify Go modules"
	@echo ""
	@echo "Frontend"
	@echo "  make web-install    Install frontend dependencies"
	@echo "  make web-build      Build frontend"
	@echo "  make web-lint       Lint frontend"
	@echo "  make web-test       Test frontend"
	@echo ""
	@echo "Utilities"
	@echo "  make check          Run format, lint and tests"
	@echo ""

# ------------------------------------------------------------------------------
# Bootstrap
# ------------------------------------------------------------------------------

.PHONY: bootstrap

bootstrap:
	@$(SCRIPT_DIR)/bootstrap.sh

# ------------------------------------------------------------------------------
# Development
# ------------------------------------------------------------------------------

.PHONY: generate

generate:
	@$(SCRIPT_DIR)/generate.sh

.PHONY: format

format:
	@$(SCRIPT_DIR)/format.sh

.PHONY: lint

lint:
	@$(SCRIPT_DIR)/lint.sh

.PHONY: test

test:
	@$(SCRIPT_DIR)/test.sh

.PHONY: build

build:
	@$(SCRIPT_DIR)/build.sh

.PHONY: dev

dev:
	@$(SCRIPT_DIR)/dev.sh

.PHONY: clean

clean:
	@$(SCRIPT_DIR)/clean.sh

# ------------------------------------------------------------------------------
# Release
# ------------------------------------------------------------------------------

.PHONY: release

release:
	@$(SCRIPT_DIR)/release.sh

# ------------------------------------------------------------------------------
# Go
# ------------------------------------------------------------------------------

.PHONY: tidy

tidy:
	go mod tidy

.PHONY: download

download:
	go mod download

.PHONY: verify

verify:
	go mod verify

# ------------------------------------------------------------------------------
# Frontend
# ------------------------------------------------------------------------------

.PHONY: web-install

web-install:
	cd web && npm install

.PHONY: web-build

web-build:
	cd web && npm run build

.PHONY: web-lint

web-lint:
	cd web && npm run lint

.PHONY: web-test

web-test:
	cd web && npm test

# ------------------------------------------------------------------------------
# Composite Targets
# ------------------------------------------------------------------------------

.PHONY: check

check: format lint test

# ------------------------------------------------------------------------------
# End
# ------------------------------------------------------------------------------