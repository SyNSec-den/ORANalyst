import os
import datetime
from collections import defaultdict

# Replace 'your_directory_path_here' with the path of your directory
directory_path = 'corpus_e2ap_final'

# Initialize a dictionary to hold counts of files for each 10-minute interval
file_counts = defaultdict(int)

# Function to convert creation time to a rounded 10-minute interval
def round_time(dt, delta=datetime.timedelta(minutes=120)):
    return datetime.datetime.min + ((dt - datetime.datetime.min) // delta) * delta

# Iterate over files in the directory
for filename in os.listdir(directory_path):
    filepath = os.path.join(directory_path, filename)
    if os.path.isfile(filepath):
        # Get the creation time of the file
        creation_time = os.path.getctime(filepath)
        # Convert to datetime
        creation_datetime = datetime.datetime.fromtimestamp(creation_time)
        # Round to nearest 10 minutes
        rounded_time = round_time(creation_datetime)
        # Increment the count for this time interval
        file_counts[rounded_time] += 1

# Print the results
print(len(file_counts))
for interval, count in sorted(file_counts.items()):
    print(f"{interval.strftime('%Y-%m-%d %H:%M')} - {count} files")
