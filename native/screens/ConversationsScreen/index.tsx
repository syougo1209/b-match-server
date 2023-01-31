import { Button } from 'react-native'
import { useNavigation} from '@react-navigation/native'
import { NativeStackNavigationProp} from '@react-navigation/native-stack';
import { Screens } from '../../Screens'
import { BottomMenu } from '../../components/elements/BottomMenu';

export const ConversationsScreen = () => {
  const navigation = useNavigation<NativeStackNavigationProp<Screens, 'Profile'>>()
  return (
    <>
      <Button
        title="Go ton"
        onPress={() =>
          navigation.navigate('Profile')
        }
      />
    </>
  );
};
