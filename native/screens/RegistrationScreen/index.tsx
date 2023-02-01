import { Button } from 'react-native'
import { useNavigation} from '@react-navigation/native'
import { NativeStackNavigationProp} from '@react-navigation/native-stack';
import { Screens } from '../../Screens'
import { MaterialCommunityIcons } from '@expo/vector-icons'
import { Container } from 'native-base'

export const RegistrationScreen = () => {
  const navigation = useNavigation<NativeStackNavigationProp<Screens, 'Conversations'>>()

  return (
    <Container safeArea>
      <MaterialCommunityIcons name="chevron-left" size={45} onPress={()=>navigation.navigate('Conversations')} />
    </Container>
  );
};
