import './index.css'
import Typewriter from 'typewriter-effect'

function Home() {
    return(
       <div className="init">
           <Typewriter
  
    options={{
        strings: ['Hacking', 'Cracking'], 
        autoStart: true,
        loop: true,
  }}
/>
       </div>
    )
}

export default Home;
