import subprocess
from datetime import datetime
import time

def run_command_and_log(command):
    # Define the command to run
    
    
    # Execute the command
    process = subprocess.Popen(command, shell=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
    stdout, stderr = process.communicate()
    
    # Check for errors
    if process.returncode != 0:
        print(f"Error running command: {stderr.decode()}")
        return
    
    # Get the current time
    current_time = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
    
    # Path to the file where the output will be appended
    output_file_path = "radamsa_cover.txt"
    
    # Append the output and the current time to the file
    with open(output_file_path, "a") as file:
        file.write(f"Time: {current_time}\n{stdout.decode()}\n")
    print(f"{current_time} successfully wrote to {output_file_path}")

if __name__ == "__main__":
    while True:
        run_command_and_log(command = "kubectl logs -n ricplt deployment-ricplt-e2term-alpha-65b4fd79fb-d4hmm | grep 'edge cover' | tail -n 5")
        run_command_and_log(command = "kubectl logs -n ricxapp ricxapp-kpimon-go-7744d5f76b-t75sc | grep 'edgeCover' | tail -n 5")
        time.sleep(600)