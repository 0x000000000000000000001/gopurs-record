module Record.Unsafe.Union where

import Data.Function.Uncurried (Fn2, runFn2)
import Unsafe.Coerce (unsafeCoerce)

foreign import _unsafeUnionFn :: forall a b c. Fn2 a b c

unsafeUnionFn :: forall r1 r2 r3. Fn2 (Record r1) (Record r2) (Record r3)
unsafeUnionFn = unsafeCoerce _unsafeUnionFn

unsafeUnion :: forall r1 r2 r3. Record r1 -> Record r2 -> Record r3
unsafeUnion = runFn2 unsafeUnionFn
