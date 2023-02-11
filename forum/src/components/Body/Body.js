import React,{useEffect, useState} from "react";
import InputForm from "../Input-Form/Input-Form";
import './Body.css'
import Post from "../Post/Post";
const Body = (props) => {  

// props.posts ? console.log("exists"): console.log("no");
// const handlePosts =()=>{
//   if (props.posts){
//     props.posts.map(({PostId, Title, Content, CreationDate, Author})=>{
//       return (<Post key={PostId} title={Title} content={Content} date={CreationDate} author={Author}/>) 
//     })
//   }
//   return (<p>asd</p>)
// }

return (
    <div className="body-container">
      <div className="body-posts-container">
      {props.createPost && <InputForm username= {props.username}/>}    
      {!props.createPost && props.posts && (props.posts.map(({PostId, Title, Content, CreationDate, Author})=>{return (<Post key={PostId} title={Title} content={Content} date={CreationDate} author={Author}/>)}))}
      </div>
    </div>
  )
};

export default Body;