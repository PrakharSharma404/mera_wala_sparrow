import { useState } from 'react'
import './App.css'

function App() {
  const [services] = useState([
    { name: 'repo-service', port: 8080, metricsPort: 2112, status: 'Active' },
    { name: 'container-service', port: 8081, metricsPort: 2112, status: 'Active' },
    { name: 'deploy-service', port: 8002, metricsPort: 2112, status: 'Active' },
  ])

  return (
    <div className="dashboard">
      <h1>Sparrow VPS Monitoring (Dummy)</h1>
      <div className="card-container">
        {services.map((svc) => (
          <div key={svc.name} className="card">
            <h2>{svc.name}</h2>
            <div className="status">
              <span className={`indicator ${svc.status.toLowerCase()}`}></span>
              {svc.status}
            </div>
            <div className="details">
              <p>API Port: {svc.port}</p>
              <p>Metrics Port: {svc.metricsPort}</p>
            </div>
          </div>
        ))}
      </div>
      <p className="note">Note: This is a dummy frontend for visual confirmation. Services run on the backend.</p>
    </div>
  )
}

export default App
