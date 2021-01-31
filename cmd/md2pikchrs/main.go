package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gebv/pikchr"
	"github.com/gebv/pikchr/markdown"
	"github.com/gebv/pikchr/markdown/syntax"
)

var (
	version = "dev"
	commit  = ""
	date    = ""
	builtBy = ""
)

var fOutDir = flag.String("out", "./out", "Sets dir for exported diagram files.")
var fInFile = flag.String("in", "./*.md", "Input markdown file.")
var fDebug = flag.Bool("debug", false, "Debug mode.")
var fVersion = flag.Bool("v", false, "Print version.")

func main() {
	log.Println("md2pikchrs version:", version+"#"+commit)

	flag.Parse()

	validFlags()
	applyFlags()

	err := filepath.Walk(*fInFile, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info == nil || info.IsDir() {
			log.Println("Skipped dir:", path)
			return nil
		}
		fileDat, err := ioutil.ReadFile(path)
		if err != nil {
			log.Println("Failed open and read file:", path, err)
			return nil
		}
		md, err := markdown.Parse(string(fileDat))
		if err != nil {
			log.Println("Failed parse md file:", path, err)
			return nil
		}
		interestedBlocks := []markdown.MarkdownCodeBlock{}
		log.Println(path, "total", len(md.CodeBlocks()), "code blocks")
		for _, block := range md.CodeBlocks() {
			if block.Language() == "pikchr" && block.StringInfoAfterLanguageName() != "" {
				interestedBlocks = append(interestedBlocks, block)
			}
		}
		log.Println(path, len(interestedBlocks), "interesting code blocks")
		for _, block := range interestedBlocks {
			fileName := strings.ReplaceAll(strings.TrimSpace(strings.ToLower(block.StringInfoAfterLanguageName())), " ", "_")
			log.Println(path, "\t", fileName, "rendering...")
			renderRes, ok := pikchr.Render(block.Content().Raw())
			if !ok {
				log.Println(path, "\t", fileName, "failed render:", renderRes.Data)
				continue
			}
			if !strings.HasSuffix(fileName, ".svg") {
				fileName += ".svg"
			}
			if err := ioutil.WriteFile(filepath.Join(*fOutDir, fileName), []byte(renderRes.Data), 0644); err != nil {
				log.Println(path, "\t", fileName, "failed write the generated file:", err)
			}
			log.Println(path, "\t", fileName, "- OK")
		}
		return nil
	})
	if err != nil {
		log.Println("failed scan dir", *fInFile, ":", err)
	}
}

func validFlags() {
	if fOutDir == nil || *fOutDir == "" {
		log.Println("Please sets the out dir.")
		os.Exit(1)
	}
	if fInFile == nil || *fInFile == "" {
		log.Println("Please sets in file\\files.")
		os.Exit(1)
	}
}

func applyFlags() {
	if fDebug != nil && *fDebug {
		log.Println("md2pikchrs enabled debug mode")
		syntax.Debug()
		syntax.ErrorVerbose()
	}

	if fVersion != nil && *fVersion {
		log.Println("md2pikchrs version:", version)
		log.Println("md2pikchrs commit:", commit)
		log.Println("md2pikchrs build date:", date)
		log.Println("md2pikchrs build by:", builtBy)
		os.Exit(0)
	}
}
