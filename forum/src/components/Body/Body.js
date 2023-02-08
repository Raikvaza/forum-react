import React from "react";
import InputForm from "../Input-Form/Input-Form";
import './Body.css'
const Body = (props) => {
  
  
return (
    <div className="body-container">
    <InputForm name= {props.name}/>
    </div>
  )


};

export default Body;