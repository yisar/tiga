import Tiga from 'tiga.js/core'

export default function IndexPage() {
  function gotoAbout() {
    Tiga.navigateTo({ url: '/about' })
  }
  return (
    <View>
      Hello World.
      <Text onTap={gotoAbout}>About</Text>
    </View>
  )
}
