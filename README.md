# vision-search

A simple CLI tool to detect web entities in images using Google Cloud Vision API.

## Features

- Detects web entities, matching images, and pages with matching images from a given image file.
- Outputs descriptions and URLs of detected entities and images.
- Easy to use from the command line.

## Requirements

- Go 1.17 or later
- Google Cloud account with Vision API enabled
- Application Default Credentials set up (see [Google Cloud documentation](https://cloud.google.com/docs/authentication/getting-started))

## Installation

### Build from source

```sh
git clone https://github.com/yourusername/webdetect.git
cd vision-search
make install
```

## Usage

```sh
vision-search <image_path>
```

- `<image_path>`: Path to the image file you want to analyze.

### Example

```sh
vision-search ./sample.jpg
```

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.
