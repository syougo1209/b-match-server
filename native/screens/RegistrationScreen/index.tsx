import { Button } from 'react-native'
import { useNavigation} from '@react-navigation/native'
import { NativeStackNavigationProp} from '@react-navigation/native-stack';
import { Screens } from '../../Screens'
import { SafeAreaContainer } from '../../components/elements/SafeAreaContainer';
import { MaterialCommunityIcons } from '@expo/vector-icons'

export const RegistrationScreen = () => {
  const navigation = useNavigation<NativeStackNavigationProp<Screens, 'Conversations'>>()

  return (
    <SafeAreaContainer>
      <MaterialCommunityIcons name="chevron-left" size={45} onPress={()=>navigation.navigate('Conversations')} />
    </SafeAreaContainer>
  );
};
