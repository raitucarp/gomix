package styles

func TransitionNormal() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(transitionBehaviorProp): "normal",
		}
	}
}

func TransitionDiscrete() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(transitionBehaviorProp): "allow-discrete",
		}
	}
}
