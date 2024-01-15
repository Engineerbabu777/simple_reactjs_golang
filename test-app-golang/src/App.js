import logo from './logo.svg'
import './App.css'
import { useEffect } from 'react'

function App () {
  const handleClick = async () => {
    fetch('http://localhost:8080/api/data')
      .then(response => response.json())
      .then(data => {
        console.log({ data })
      })
  }
  return (
    <div className='App'>
      <button onClick={handleClick}>Click me</button>
    </div>
  )
}

export default App
