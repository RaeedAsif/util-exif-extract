# Image Metadata Extractor

This is a command-line tool written in Go that extracts GPS metadata from images and saves the data in a CSV file. The tool supports the following image formats: ".jpg", ".jpeg", ".png", ".gif", ".bmp", ".svg".

## How to Run

To run the tool, you need to have Go installed on your machine. Then, follow these steps:

1. Clone this repository:

```
git clone https://github.com/your-username/image-metadata-extractor.git
```

2. Run the tool using the following command:

```
make run
```

This will run the tool and generate two CSV files: `result.csv` and `error_result.csv`. The `result.csv` file contains the GPS metadata for images that have GPS coordinates, while the `error_result.csv` file contains a list of images that do not have GPS metadata or could not be processed.

4. View the results in a web browser by accessing the following URLs:
- http://localhost:8080/ - displays index html to navigate to csv's
- http://localhost:8080/result - displays the GPS metadata for the images in `result.csv`
- http://localhost:8080/error - displays the list of images in `error_result.csv` that could not be processed
