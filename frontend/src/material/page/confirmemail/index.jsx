import './index.css'
import { useForm } from 'react-hook-form'
import { yupResolver } from '@hookform/resolvers/yup'
import * as Yup from 'yup'
import axios from "axios"
import Cookies from 'universal-cookie';

export default function Confirmotp() {
  const cookies = new Cookies();
  //
  var dff= axios({url:'http://localhost:9000/q',method:"post",headers:{"jwt":cookies.get("Destroy")}});
    const formSchema = Yup.object().shape({
        otp: Yup.string()
          .required('otp is require')
          .matches(1-9, 'otp is invalid')
          .min(6, 'otp must be 6 digit')
          .max(6, 'otp must be 6 digit'),
       
      })
      
      const formOptions = { resolver: yupResolver(formSchema) }
          const { register, handleSubmit, formState } = useForm(formOptions)
          const { errors } = formState
          async function onSubmit(data) {
        
            console.log(data)
          }

    return(
      
        <div className="wrapper">
          
               <form  onSubmit={handleSubmit(onSubmit)} className="otpformad">
                   <div className='field'>
                       <label className='otpla'>otp</label>
                     <input name="otp" type="password" placeholder='******' {...register('otp')} maxLength="6" minLength="6" className={` ${errors.otp ? 'is-invalid' : ''}`}/>
                     <div className="invalid-feedback red">{errors.otp?.message}</div>

                   </div>
                   <div className="field btn">
                            <div className="btn-layer" />
                            <input type="submit" value="Confirm" />
                        </div>
               </form>
          
        </div>

    )
};