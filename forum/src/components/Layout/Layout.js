import Header from "../Header/Header";
import Body from "../Body/Body";
function Layout({ children }) {
  return (
    <>
      <Header/>
      <Body>
        {children}
      </Body>
    </>
  );
}

export default Layout;
