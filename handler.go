package main

import (
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "net/http"
    "io/ioutil"
)

// ConvertToAPK converts a given website URL to an APK file.
func ConvertToAPK(url string, outputDir string) error {
    // Create a temporary directory for storing the APK build files
    tmpDir, err := os.MkdirTemp(outputDir, "apk-builder-")
    if err != nil {
        return fmt.Errorf("failed to create temp directory: %w", err)
    }
    defer os.RemoveAll(tmpDir) // Clean up

    // Use a tool like webview to build the APK
    cmd := exec.Command("webview-build-tool", url, tmpDir)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to build APK: %w", err)
    }

    // Move the APK to the desired output location
    apkPath := filepath.Join(tmpDir, "output.apk")
    if err := os.Rename(apkPath, filepath.Join(outputDir, "output.apk")); err != nil {
        return fmt.Errorf("failed to move APK: %w", err)
    }

    return nil
}

// A simple HTTP handler for triggering APK conversion
func APKHandler(w http.ResponseWriter, r *http.Request) {
    url := r.URL.Query().Get("url")
    if url == "" {
        http.Error(w, "URL is required", http.StatusBadRequest)
        return
    }
    outputDir := "./output"
    if err := ConvertToAPK(url, outputDir); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    fmt.Fprintf(w, "APK file created successfully!")
}