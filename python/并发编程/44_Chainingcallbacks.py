

with ThreadPoolExecutor(max_workers=3) as executor:
  future = executor.submit(task, (2))
  future.add_done_callback(taskDone)
  future.add_done_callback(secondTaskDone)