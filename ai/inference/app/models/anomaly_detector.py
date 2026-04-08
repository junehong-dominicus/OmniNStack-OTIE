from .base import BaseThreatModel
import numpy as np

class AnomalyDetector(BaseThreatModel):
    def __init__(self, threshold=0.85):
        self.threshold = threshold

    def predict(self, features: np.ndarray) -> dict:
        # Mock logic: if sum/mean exceeds threshold, mark as threat
        score = np.mean(features)
        is_threat = bool(score > self.threshold)
        return {
            'score': float(score),
            'is_threat': is_threat,
            'model': 'anomaly_v1'
        }

