import requests
import zipfile
import io
import os
import sys

if __name__ != "__main__":
    print("This script is not intended to be imported as module")
    exit(1)

if len(sys.argv) != 2:
    print("Usage: python download_data.py <data_path>")
    exit(1)
print("Creating data directory")
data_path = sys.argv[1]
if not os.path.exists(data_path):
    os.makedirs(data_path)

print(f"Downloading dataset to {data_path}")
uri = "https://www.kaggle.com/api/v1/datasets/download/shashwatwork/web-page-phishing-detection-dataset"
response = requests.get(uri)
buffer = io.BytesIO(response.content)
print("Download complete")
print("Extracting dataset")
zipfile.ZipFile(buffer).extractall(data_path)
print("Extraction complete")
print("Dataset downloaded and extracted successfully")