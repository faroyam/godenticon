# Godenticon

Godenticon is a tiny package that generates identicons based on the hash of provided and it is inspired by [Pydenticon](https://github.com/azaghal/pydenticon)

## Example

```go
// Calculate MD5 checksum of the data
hash := md5.Sum([]byte("Gopher!"))

// Create an image using hash and slice of colours 
img, _ := g.GenarateImage(hash, g.ColorGopher)

// Create a new file to be written down a specified path 
f, _ := os.Create("image.png")

// Write image into a file
g.SaveImage(img, f)
```

It will output image:

![Gopher!](example/image.png?raw=true)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
