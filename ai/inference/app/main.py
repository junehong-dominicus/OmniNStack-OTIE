from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
from typing import List
import numpy as np
from app.models.anomaly_detector import AnomalyDetector

app = FastAPI(title='OTIE Precision AI Engine', version='1.0.0')
model = AnomalyDetector()

class InferenceRequest(BaseModel):
    event_id: str
    features: List[float]

class InferenceResponse(BaseModel):
    event_id: str
    score: float
    is_threat: bool
    description: str

@app.get('/health')
def health():
    return {'status': 'up', 'engine': 'precision-ai-v1'}

@app.post('/predict', response_model=InferenceResponse)
def predict(request: InferenceRequest):
    try:
        data = np.array(request.features)
        result = model.predict(data)
        return {
            'event_id': request.event_id,
            'score': result['score'],
            'is_threat': result['is_threat'],
            'description': f'Analysis complete using {result["model"]}'
        }
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))

