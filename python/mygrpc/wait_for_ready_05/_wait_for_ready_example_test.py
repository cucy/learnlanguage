import wait_for_ready_example

import unittest
import logging


class WaitForReadyExampleTest(unittest.TestCase):

    def test_wait_for_ready_example(self):
        wait_for_ready_example.main()
        # No unhandled exception raised, no deadlock, test passed!


if __name__ == '__main__':
    logging.basicConfig()
    unittest.main(verbosity=2)
