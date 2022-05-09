/* eslint-disable jsx-a11y/anchor-is-valid */
import { useTranslation } from "react-i18next";



import { Link } from 'react-router-dom'
import './sign.css'




         

    

function RegisLog() {
  const { t } = useTranslation(["login"]);
    return (
      
                
                <div className="wrapper">
  <div className="title-text">
    <div className="title login">
      {t("login")}
    </div>
    
  </div>
  <div className="form-container">
    <div className="form-inner">
      <form action="#" className="login">
        <div className="field">
          <input type="text" placeholder="Email Address" required />
        </div>
        <div className="field">
          <input type="password" placeholder="Password" required />
        </div>
        <div className="pass-link">
          <a href="#" style={{color: "#9FEF00"}}>Forgot password?</a>
        </div>
        <div className="field btn">
          <div className="btn-layer" />
          <input type="submit" value="Login"/>
        </div>
        <div className="signup-link">
          Not a member? <Link to="/signup"><a>Signup now</a></Link>
        </div>
      </form>
      
    </div>
  </div>
</div>

          
    )
}

export default RegisLog;