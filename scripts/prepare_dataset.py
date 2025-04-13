import pandas as pd
import os

if __name__ != "__main__":
    print("This script is not intended to be imported as module")
    exit(1)

if len(os.sys.argv) != 2:
    print("Usage: python prepare_dataset.py <data_path>")
    exit(1)

data_path = os.sys.argv[1]
input_path = os.path.join(data_path, "dataset_phishing.csv")
ds = pd.read_csv(input_path)
print("Original dataset:")
print(ds.head())
print("Columns:")
print(ds.columns)
print("Data types:")
print(ds.dtypes)

labels = ds["status"].unique()
print("Labels:")
print(labels)
urls_grouped = ds.groupby("status")["url"].apply(list)
phishing_urls = urls_grouped["phishing"]
legitimate_urls = urls_grouped["legitimate"]
print("Phishing URLs:")
print(phishing_urls[:10])
print("Legitimate URLs:")
print(legitimate_urls[:10])

print("Number of phishing URLs:")
print(len(phishing_urls))
print("Number of legitimate URLs:")
print(len(legitimate_urls))

print("Saving to output path...")
phishing_path = os.path.join(data_path, "phishing_urls.csv")
phishing_urls = pd.DataFrame(phishing_urls)
phishing_urls.to_csv(phishing_path, index=False, header=False)

legitimate_path = os.path.join(data_path, "legitimate_urls.csv")
legitimate_urls = pd.DataFrame(legitimate_urls)
legitimate_urls.to_csv(legitimate_path, index=False, header=False)


print("Output paths:")
print(phishing_path)
print(legitimate_path)