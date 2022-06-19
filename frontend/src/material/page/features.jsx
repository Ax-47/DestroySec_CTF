import axios from "axios"
import Cookies from 'universal-cookie';

 function Features() {
    (async function(){
        const cookies = new Cookies();
        var res = axios({url:'http://localhost:9000/checkpermis',method:"post",headers:{"jwt":cookies.get("Destroy"),"X-API-KEY":"ax47"},data:{"jwt":cookies.get("Destroy")}});
        res = await res
        var check
        if (res.status ===200){
            check="logined"
        }else{
            check="404 not found"
    
        }
    })()
    return (
        <div>check</div>
    )
}

export default Features;