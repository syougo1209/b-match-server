import { Button } from 'react-native'
import { useNavigation} from '@react-navigation/native'
import { NativeStackNavigationProp} from '@react-navigation/native-stack';
import { Screens } from '../../Screens'
import { BottomMenu } from '../../components/elements/BottomMenu';

export const ProfileScreen = () => {
  const navigation = useNavigation<NativeStackNavigationProp<Screens, 'Conversations'>>()
  return (
    <>
      <Button
        title="Go n"
        onPress={() =>
          navigation.navigate('Conversations')
        }
      />
    </>
  );
};
