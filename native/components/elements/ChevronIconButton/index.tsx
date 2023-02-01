import { Icon, IconButton } from 'native-base'
import { MaterialCommunityIcons } from '@expo/vector-icons'
import { FC } from 'react'

type ChevronIconButton = {
  onPress: ()=> void
}

export const ChevronIconButton: FC<ChevronIconButton> = (props) => {
  const { onPress } = props
  return <IconButton icon={<Icon as={MaterialCommunityIcons} size="50" name="chevron-left" color="black"/>} onPress={onPress}/>
}
