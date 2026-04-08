from abc import ABC, abstractmethod
import numpy as np

class BaseThreatModel(ABC):
    @abstractmethod
    def predict(self, features: np.ndarray) -> dict:
        pass

