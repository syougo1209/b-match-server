import { Button } from 'react-native'
import { useNavigation} from '@react-navigation/native'
import { NativeStackNavigationProp} from '@react-navigation/native-stack';
import { Screens } from '../../Screens'
import { SafeAreaContainer } from '../../components/elements/SafeAreaContainer';

export const ConversationsScreen = () => {
  const navigation = useNavigation<NativeStackNavigationProp<Screens, 'Profile'>>()
  return (
    <SafeAreaContainer>
      <Button
        title="Go ton"
        onPress={() =>
          navigation.navigate('Profile')
        }
      />
    </SafeAreaContainer>
  );
};
