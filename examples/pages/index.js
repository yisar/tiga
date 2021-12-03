import Tiga from '@tiga/core'

export default function App() {
  function gotoAbout() {
    Tiga.navigateTo('/about')
  }
  return (
    <View>
      Hello World.
      <Text onTap={gotoAbout}>About</Text>
    </View>
  )
}
