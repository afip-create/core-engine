# core-engine/main.py

import os
import logging
from typing import Dict

# Set up logging configuration
logging.basicConfig(
    format='%(asctime)s [%(levelname)s] %(message)s',
    level=logging.INFO
)

def get_config() -> Dict:
    """Load configuration from environment variables"""
    config: Dict = {
        'database': {
            'host': os.environ.get('DATABASE_HOST'),
            'port': int(os.environ.get('DATABASE_PORT')),
            'username': os.environ.get('DATABASE_USERNAME'),
            'password': os.environ.get('DATABASE_PASSWORD'),
            'database': os.environ.get('DATABASE_NAME')
        }
    }
    return config

def main():
    config = get_config()
    logging.info('Loaded configuration: %s', config)

if __name__ == '__main__':
    main()